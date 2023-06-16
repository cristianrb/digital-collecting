CREATE TABLE items
(
    id   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR NOT NULL UNIQUE,
    description VARCHAR NOT NULL,
    image VARCHAR,
    price FLOAT NOT NULL
);

CREATE TABLE user_items
(
    user_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,

    PRIMARY KEY (user_id, item_id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES users (id)
);