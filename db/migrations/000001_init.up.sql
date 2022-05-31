CREATE TABLE IF NOT EXISTS users
(
    id      integer primary key,
    phone   varchar(12),
    name    varchar(50) not null,
    surname varchar(50) not null
);

CREATE TABLE IF NOT EXISTS doctors
(
    id      serial primary key,
    name    varchar(50) not null,
    surname varchar(50) not null
);

CREATE TABLE IF NOT EXISTS slots
(
    id        serial primary key,
    doctor_id serial references doctors (id) ON DELETE CASCADE,
    slot      timestamp not null,
    UNIQUE (doctor_id, slot)
);

CREATE TABLE IF NOT EXISTS appointments
(
    id        serial primary key,
    doctor_id serial references doctors (id) ON DELETE CASCADE,
    user_id   serial references users (id) ON DELETE CASCADE,
    slot      timestamp,
    UNIQUE (doctor_id, user_id, slot)
);

