create table users
(
    id                int not null auto_increment,
    email             varchar(255),
    password          varchar(255),
    username          varchar(255),
    created_at        timestamp null,
    updated_at        timestamp null,
    email_verified_at timestamp null,
    primary key (id)
)