
Table category {
  category_id int [pk]
  name varchar
  description text
  create_at timestamp
  update_at timestamp
  delete_at timestamp
}

Table inventory_items {
  inventory_items_id int [pk]
  category_id int 
  name varchar
  price int
  purchase_date timestamp
  create_at timestamp
  update_at timestamp
  delete_at timestamp
}




Ref: "category"."category_id" < "inventory_items"."category_id"