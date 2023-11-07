# Скрипты для БД

CREATE TABLE IF NOT EXISTS "Category"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "activeFlg" BIGINT NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    "rang" BIGINT NOT NULL DEFAULT 0
);
CREATE index if not EXISTS "rang_index" ON "Category" (
"rang"
);

CREATE TABLE IF NOT EXISTS "Location"(
    "id" SERIAL PRIMARY KEY,
    "activeFlg" BIGINT NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    "locationFilePath" VARCHAR(511) NOT NULL
);

CREATE TABLE IF NOT EXISTS "Items"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "activeFlg" BIGINT NOT NULL DEFAULT 0,
    "createdAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    "rang" BIGINT NOT NULL DEFAULT 0,
    "logoImg" VARCHAR(255) NOT NULL DEFAULT "logomock",
    "screenshotsImg" VARCHAR(255) NOT NULL DEFAULT "imgmock",
    "description" VARCHAR(511) NOT NULL,
    "locationId" BIGINT REFERENCES "Location" (id),
    "status" BIGINT NOT NULL DEFAULT 0
);

ALTER TABLE Items
ADD COLUMN "logoImg" VARCHAR(255) NOT NULL DEFAULT "logomock",
ADD COLUMN "ScreenshotsImg" VARCHAR(255) NOT NULL DEFAULT "imgmock",
ADD COLUMN "Status" int DEFAULT 0;

CREATE index if not EXISTS "locationId_index" ON "Items" (
"locationId"
);

CREATE index if not EXISTS "rang_index" ON "Items" (
"rang"
);

CREATE index if not EXISTS "active_index" ON "Items" (
"activeFlg"
);

CREATE index if not EXISTS "status_index" ON "Items" (
"status"
);

Category_items
CREATE TABLE IF NOT EXISTS "Category_items"(
    "id" BIGINT REFERENCES "Category" (id),
    "itemId" BIGINT REFERENCES "Items" (id)
);

CREATE index if not EXISTS "id_index" ON "Category_items" (
"id"
);

CREATE index if not EXISTS "Item_index" ON "Category_items" (
"ItemId"
);


CREATE MATERIALIZED VIEW IF NOT EXISTS mvcategory
    AS
    select c.id as catid , c.title as cattitle, c.rang as catrang , 
    i.id as itemid, i.title as itemtitle, i."logoImg"  as itemlogo , row_number () OVER (
            PARTITION BY c.id
            ORDER BY i.rang asc
        ) as group_nn
from "Category" as c 
join "Category_items" as ci
on c.id = ci."id"
join "Items" as i 
on i.id = ci."itemId" 
where c."activeFlg" = 1
and i."activeFlg"  = 1
order by c.rang , i.rang asc;

create UNIQUE  index if not EXISTS "uniq_index" ON "mvcategory" (
"catid" , "group_nn"
);
create index if not EXISTS "catId_index" ON "mvcategory" (
"catid"
);

create UNIQUE  index if not EXISTS "uniqLink_index" ON "Category_items" (
"id" , "itemId"
);

