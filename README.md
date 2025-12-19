# 🚀 Office Inventory System (CLI)

> An Efficient Terminal-Based Solution for Managing Office Inventory Assets and Financial Reporting.

[![Go](https://img.shields.io/badge/Language-Go%201.25.5-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?style=for-the-badge&logo=postgresql)](https://www.postgresql.org/)
[![CLI](https://img.shields.io/badge/Framework-Cobra%20CLI-555555?style=for-the-badge&logo=cobra)](https://github.com/spf13/cobra)

This CLI utility is designed to manage the lifecycle of office inventory assets, covering categorization, stock control, and essential financial analysis, including depreciation and investment value tracking.

CLI Execution Name: `./InventorySystem-CLI`

---

## ✨ Core Features

The application functionalities are structured into two main modules:

### 1. Inventory Category Management

| Command | Alias | Brief Description | Example Usage |
| :--- | :--- | :--- | :--- |
| `get-all-category` | | Displays a complete list of all categories. | `./InventorySystem-CLI get-all-category` |
| `create-category` | | Registers a new category. | `./InventorySystem-CLI create-category -n "ATK" -d "Office Supplies"` |
| `get-category-byID` | `-i` | Views specific category details by ID. | `./InventorySystem-CLI get-category-byID -i 5` |
| `updated-category-byID`| `-i` | Modifies the name or description of an existing category. | `./InventorySystem-CLI updated-category-byID -i 5 -n "New Tools"` |
| `delete-category` | `-i` | Permanently removes a category. | `./InventorySystem-CLI delete-category -i 5` |

### 2. Asset Reporting & Analysis

| Command | Alias | Brief Description | Example Usage |
| :--- | :--- | :--- | :--- |
| `items-morethan-100days`| | **Replacement Report:** Lists items used for more than 100 days (suggesting replacement). | `./InventorySystem-CLI items-morethan-100days` |
| `invesment-depreciation`| `-i` | **Financial Report:** Displays the current investment value and total depreciation for an item by ID. | `./InventorySystem-CLI invesment-depreciation -i 12` |
| `find-inventory-byname` | `-n` | **Quick Search:** Finds and displays the details of an inventory item by name. | `./InventorySystem-CLI find-inventory-byname -n "Laptop X20"` |

> **💡 Depreciation Logic:** Depreciation is calculated using the declining balance method, assuming a residual value rate of 20% per year (using the factor `0.8` in the PostgreSQL queries).

---

## 🛠️ Requirements & Installation

### Prerequisites

* **Go (Golang):** Version `1.25.5` or newer.
* **PostgreSQL:** The database must be installed and running.

### Database Setup

The current database connection configuration (found in `database/database.go`) is:

| Parameter | Value |
| :--- | :--- |
| **Database Name** | `inventory` |
| **User** | `your username` |
| **Password** | `your password` |
| **Host** | `localhost` |
| **Port** | `5432` |

1.  Create a new PostgreSQL database named `inventory`.
2.  Run the following SQL script to create the necessary tables:

```sql
-- Create Category Table
CREATE TABLE category(
	category_id SERIAL PRIMARY KEY ,
	name VARCHAR(100),
	description TEXT,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp
);

-- Create Inventory Items Table
CREATE TABLE inventory_items(
	inventory_items_id SERIAL PRIMARY KEY ,
	category_id int,
	name VARCHAR(100),
	price NUMERIC(15,0),
	purchase_date DATE,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp
);
```
## video tutorial : https://drive.google.com/file/d/1Tb6PBA-Mx_ks_e3Iq6jH_Zbe_it3PsNK/view?usp=sharing