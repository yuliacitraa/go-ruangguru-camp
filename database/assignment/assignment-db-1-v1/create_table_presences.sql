create table attendaces (
    id int not null,
    user_id int not null,
    status varchar(100) not null
);

create table presences (
    id int not null,
    user_id int not null,
    presence_date date not null,
    status varchar(50) not null,
    location varchar(255),
    description varchar(255),
    image_presence varchar(255),
    image_location varchar(255)
);