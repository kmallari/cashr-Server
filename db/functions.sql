CREATE OR REPLACE FUNCTION update_categories(
    ids char(36)[],
    icons varchar(24)[],
    labels varchar(60)[],
    is_expenses bool[]
) RETURNS void AS
$$
BEGIN
    UPDATE category
    SET icon       = coalesce(icons[idx], icon),
        label      = coalesce(labels[idx], label),
        is_expense = coalesce(is_expenses[idx], is_expense)
    FROM unnest(ids, icons, labels, is_expenses) WITH ORDINALITY AS t(id, icon, label, is_expense, idx)
    WHERE category.id = t.id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION all_categories()
    RETURNS SETOF category AS
$$
SELECT *
FROM category;
$$
    LANGUAGE sql;

select * from all_categories();