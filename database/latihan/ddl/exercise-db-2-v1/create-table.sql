CREATE TABLE persons (
    id int primary key,
    nik varchar(255) not null unique,
    fullname varchar(255) not null,
    gender varchar(50) not null,
    birth_date date not null,
    is_married boolean,
    height float,
    weight float,
    address text
);
