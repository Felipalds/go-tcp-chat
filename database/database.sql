PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users(
    id          INTEGER PRIMARY KEY CHECK(id != 0),
    name        TEXT    NOT NULL UNIQUE
--     aes_key     TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id      INTEGER PRIMARY KEY,
    name    TEXT    NOT NULL UNIQUE CHECK(id != 0),
    type TEXT NOT NULL,
    password TEXT,
    admin_id   INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
--     CHECK(LENGTH(password)=32 OR (LENGTH(password)=0)),
    CHECK(type = 'PUBLICA' OR type='PRIVADA')
);

CREATE TABLE  user_room(
    room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_banned BOOLEAN NOT NULL,
    UNIQUE(room_id, user_id)
);


-- CREATE TABLE room_banned(
--     room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
--     user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
--     UNIQUE(room_id, user_id)
-- );

CREATE TABLE messages(
    id      INTEGER PRIMARY KEY CHECK(id != 0),
    msg     TEXT    NOT NULL,
    CHECK(LENGTH(msg) != 0)
);

CREATE TABLE user_messages(
    id      INTEGER PRIMARY KEY CHECK(id != 0),
    user_id INTEGER NOT NULL REFERENCES users(id)     ON DELETE CASCADE,
    message_id  INTEGER NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    UNIQUE(user_id, message_id)
);
