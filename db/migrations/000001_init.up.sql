CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users
(
    id      uuid NOT NULL DEFAULT uuid_generate_v1() primary key,
    phone   varchar(12) unique,
    name    varchar(50) not null,
    surname varchar(50) not null
);

CREATE TABLE IF NOT EXISTS doctors
(
    id      uuid NOT NULL DEFAULT uuid_generate_v1() primary key,
    name    varchar(50) not null,
    spec varchar(50) not null
);

CREATE TABLE IF NOT EXISTS slots
(
    doctor_id uuid references doctors (id) ON DELETE CASCADE,
    slot      timestamp not null,
    UNIQUE (doctor_id, slot)
);

CREATE TABLE IF NOT EXISTS appointments
(
    doctor_id uuid references doctors (id) ON DELETE CASCADE,
    user_id   uuid references users (id) ON DELETE CASCADE,
    slot      timestamp,
    FOREIGN KEY (doctor_id, slot) references slots(doctor_id, slot),
    UNIQUE (doctor_id, slot)
);

insert into doctors (name, spec) values ('Alex', 'GP') returning id;
insert into slots (doctor_id, slot) select id, '2022-06-07 16:00:00' from doctors where name='Alex' and spec='GP';
insert into slots (doctor_id, slot) select id, '2022-06-07 16:30:00' from doctors where name='Alex' and spec='GP';
insert into slots (doctor_id, slot) select id, '2022-06-07 15:00:00' from doctors where name='Alex' and spec='GP';
insert into slots (doctor_id, slot) select id, '2022-06-07 15:30:00' from doctors where name='Alex' and spec='GP';

insert into doctors (name, spec) values ('Greg', 'Surgeon') returning id;
insert into slots (doctor_id, slot) select id, '2022-06-07 14:00:00' from doctors where name='Greg' and spec='Surgeon';
insert into slots (doctor_id, slot) select id, '2022-06-07 14:30:00' from doctors where name='Greg' and spec='Surgeon';
insert into slots (doctor_id, slot) select id, '2022-06-07 15:00:00' from doctors where name='Greg' and spec='Surgeon';
insert into slots (doctor_id, slot) select id, '2022-06-07 15:30:00' from doctors where name='Greg' and spec='Surgeon';

insert into users (phone, name, surname) values ('89196951400', 'Robert', 'M.');
insert into users (phone, name, surname) values ('89197951400', 'Max', 'R.');
insert into users (phone, name, surname) values ('89198951400', 'Peter', 'X.');
insert into users (phone, name, surname) values ('89199951400', 'Rick', 'R.');