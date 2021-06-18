use users

create table if not exists users(
   id int not null auto_increment,
   name varchar(100) not null,
   primary key ( id )
);

create table if not exists tags(
   id int not null auto_increment,
   value varchar(100) not null,
   primary key ( id )   
)

create table if not exists users_tags(
   user_id int not null,
   tag_id int not null,
   foreign key (user_id) references users(id),
   foreign key (tag_id) references tags(id)
)