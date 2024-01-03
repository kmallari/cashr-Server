package auth

import (
	"encoding/json"
	"github.com/kmallari/cashr-Server/api/utils"
	"github.com/kmallari/cashr-Server/email"
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	_ "github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"net/http"
	"os"
)

/*
JS Function to generate API keys

function generateApiKey() {
  const validChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';
  let apiKey = '';
  const numOfChars = 100;
  for (let i = 0; i < numOfChars; i++) {
    const randomIndex = Math.floor(Math.random() * validChars.length);
    apiKey += validChars.charAt(randomIndex);
  }

  return apiKey;
}
*/

func init() {
	apiDomain := os.Getenv("API_DOMAIN")
	clientDomain := os.Getenv("CLIENT_DOMAIN")

	supertokensDomain := os.Getenv("SUPERTOKENS_DOMAIN")
	supertokensApiKey := os.Getenv("SUPERTOKENS_API_KEY")

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	githubClientId := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: supertokensDomain,
			APIKey:        supertokensApiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "Cashr Server",
			APIDomain:     apiDomain,
			WebsiteDomain: clientDomain,
		},
		RecipeList: []supertokens.Recipe{
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired,
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {
						*originalImplementation.SendEmail = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
							// TODO: run some logic before sending the email
							email.SendEmailVerificationEmail(input.EmailVerification.User.Email, input.EmailVerification.EmailVerifyLink)
							// TODO: run some logic post sending the email
							return nil
						}

						return originalImplementation
					},
				},
				Override: &evmodels.OverrideStruct{
					APIs: func(originalImplementation evmodels.APIInterface) evmodels.APIInterface {
						ogVerifyEmailPOST := *originalImplementation.VerifyEmailPOST
						*originalImplementation.VerifyEmailPOST = func(token string, sessionContainer sessmodels.SessionContainer, tenantId string, options evmodels.APIOptions, userContext supertokens.UserContext) (evmodels.VerifyEmailPOSTResponse, error) {
							resp, err := ogVerifyEmailPOST(token, sessionContainer, tenantId, options, userContext)
							if err != nil {
								return evmodels.VerifyEmailPOSTResponse{}, err
							}

							if resp.OK != nil {
								id := resp.OK.User.ID

								// TODO: post email verification logic
								createDefaultCategoriesForUser(id)
							}

							return resp, nil
						}
						return originalImplementation
					},
				},
			}),
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {
						*originalImplementation.SendEmail = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
							// TODO: create and send password reset email
							email.SendResetPasswordEmail(input.PasswordReset.User.Email, input.PasswordReset.PasswordResetLink)
							// Or use the original implementation which calls the default service,
							// or a service that you may have specified in the EmailDelivery object.
							//return originalSendEmail(input, userContext)
							return nil
						}

						return originalImplementation
					},
				},
				/*
				   We use different credentials for different platforms when required. For example the redirect URI for Github
				   is different for Web and mobile. In such a case we can provide multiple providers with different client Ids.

				   When the frontend makes a request and wants to use a specific clientId, it needs to send the clientId to use in the
				   request. In the absence of a clientId in the request the SDK uses the default provider, indicated by `isDefault: true`.
				   When adding multiple providers for the same type (Google, Github etc), make sure to set `isDefault: true`.
				*/
				Providers: []tpmodels.ProviderInput{
					// We have provided you with development keys which you can use for testsing.
					// IMPORTANT: Please replace them with your own OAuth keys for production use.
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientType:   "web",
									ClientID:     googleClientId,
									ClientSecret: googleClientSecret,
								},
								//{
								//	// we use this for mobile apps
								//	ClientType:   "mobile",
								//	ClientID:     "", // requires new client ID
								//	ClientSecret: "", // this is empty because we follow Authorization code grant flow via PKCE for mobile apps (Google doesn't issue a client secret for mobile apps).
								//},
							},
						},
					},
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "github",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientType:   "web",
									ClientID:     githubClientId,
									ClientSecret: githubClientSecret,
								},
								//{ // requires new client credentials
								//	// We use this for mobile apps
								//	ClientType:   "mobile",
								//	ClientID:     githubClientId,
								//	ClientSecret: githubClientSecret,
								//},
							},
						},
					},
				},
				Override: &tpepmodels.OverrideStruct{
					Functions: func(originalImplementation tpepmodels.RecipeInterface) tpepmodels.RecipeInterface {
						// create a copy of the originalImplementation
						originalThirdPartySignInUp := *originalImplementation.ThirdPartySignInUp

						// override the thirdparty sign in / up function
						*originalImplementation.ThirdPartySignInUp = func(thirdPartyID, thirdPartyUserID, email string, oAuthTokens tpmodels.TypeOAuthTokens, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignInUpResponse, error) {

							// TODO: some pre sign in / up logic

							resp, err := originalThirdPartySignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpepmodels.SignInUpResponse{}, err
							}

							if resp.OK != nil {
								userID := resp.OK.User.ID

								if resp.OK.CreatedNewUser {
									// TODO: Post sign up logic
									createDefaultCategoriesForUser(userID)
								} else {
									// TODO: Post sign in logic
								}
							}

							return resp, err
						}

						return originalImplementation
					},
				},
			}),
			session.Init(nil),
			dashboard.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}

func SessionInfo(w http.ResponseWriter, r *http.Request) {
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	if sessionContainer == nil {
		w.WriteHeader(500)
		w.Write([]byte("no session found"))
		return
	}
	sessionData, err := sessionContainer.GetSessionDataInDatabase()
	if err != nil {
		err = supertokens.ErrorHandler(err, r, w)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		return
	}
	w.WriteHeader(200)
	w.Header().Add("content-type", "application/json")
	bytes, err := json.Marshal(map[string]interface{}{
		"sessionHandle":      sessionContainer.GetHandle(),
		"userId":             sessionContainer.GetUserID(),
		"accessTokenPayload": sessionContainer.GetAccessTokenPayload(),
		"sessionData":        sessionData,
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error in converting to json"))
	} else {
		w.Write(bytes)
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	sessionContainer := session.GetSessionFromRequestContext(r.Context())
	if sessionContainer == nil {
		w.WriteHeader(500)
		w.Write([]byte("no session found"))
		return
	}

	userID := sessionContainer.GetUserID()

	// You can learn more about the `User` object over here https://github.com/supertokens/core-driver-interface/wiki
	userInfo, err := thirdpartyemailpassword.GetUserById(userID)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, userInfo)
}
