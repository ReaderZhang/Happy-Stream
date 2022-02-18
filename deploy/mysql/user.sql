create table permission
(
    id        int auto_increment
        primary key,
    parent_id varchar(100) null,
    url       varchar(255) null
);

create table role
(
    id        int auto_increment
        primary key,
    role_name varchar(100)  null,
    pid       int default 1 null,
    constraint role_FK
        foreign key (pid) references permission (id)
);

create table member
(
    id         int auto_increment
        primary key,
    username   varchar(16)                                                                                                                                                                                                                                                                                            not null,
    phone      varchar(11)                                                                                                                                                                                                                                                                                            not null,
    password   varchar(25)                                                                                                                                                                                                                                                                                            not null,
    email      varchar(100)                                                                                                                                                                                                                                                                                           null,
    avatar     varchar(555)   default 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F0d%2F96%2F7b%2F0d967b771102473f963e34bd916e5dc0.jpeg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1647714492&t=0a9a941aede934f1f564b895897a0d52' null,
    recharge   decimal(16, 2) default 0.00                                                                                                                                                                                                                                                                            null,
    balance    decimal(16, 2) default 0.00                                                                                                                                                                                                                                                                            null,
    created_at date                                                                                                                                                                                                                                                                                                   null,
    updated_at date                                                                                                                                                                                                                                                                                                   null,
    deleted_at date                                                                                                                                                                                                                                                                                                   null,
    rid        int                                                                                                                                                                                                                                                                                                    null,
    constraint member_FK
        foreign key (rid) references role (id)
);

