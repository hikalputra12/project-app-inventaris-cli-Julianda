CREATE TABLE category(
	category_id SERIAL PRIMARY KEY ,
	name VARCHAR(100),
	description TEXT,
	created_at timestamp default current_timestamp,
)
CREATE TABLE inventory_items(
	inventory_items_id SERIAL PRIMARY KEY ,
	category_id int,
	name VARCHAR(100),
	price NUMERIC(15,0),
	purchase_date DATE,
	created_at timestamp default current_timestamp,
)

--fitur kategori barang

--menampilkan daftar category--
SELECT category_id, name, description FROM category

--menambah category baru--
INSERT INTO category (name,description) VALUES
($1, $2) RETURNING id

--melihat detail kategori berdarsarkan id
SELECT name,description FROM category WHERE category_id=$1;


--mengedit category berdasarkan id
UPDATE category
SET name=$1, description=$2
WHERE category_id=$3

--menghapus category berdarsarkan id
DELETE FROM category 
WHERE category_id=$1;

--Menampilkan barang yang sudah digunakan lebih dari 100 hari dan perlu diganti.
SELECT name,
	   price,
       CURRENT_DATE -purchase_date AS total_usage_days
FROM inventory_items
WHERE CURRENT_DATE- purchase_date >100;

--Laporan Investasi dan Depresiasi

--Menampilkan total nilai investasi seluruh barang setelah perhitungan depresiasi sebesar 20% pertahun.
SELECT
    SUM(price * POWER(
        0.8,
        DATE_PART('year', AGE(CURRENT_DATE, purchase_date))
    )) AS total_investment_value
FROM inventory_items

--Menampilkan nilai investasi dan depresiasi untuk barang tertentu berdasarkan ID.
SELECT
    name,
    price AS initial_price,

    price * POWER(
        0.8,
        DATE_PART('year', AGE(CURRENT_DATE, purchase_date))
    ) AS investment_value,

    price - (
        price * POWER(
            0.8,
            DATE_PART('year', AGE(CURRENT_DATE, purchase_date))
        )
    ) AS depreciation

FROM inventory_items
WHERE inventory_items_id = $1;
