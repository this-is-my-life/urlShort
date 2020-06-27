create user 'urlShort'@'localhost' identified by 'urlshort';
grant all privileges on urlShort.* to 'urlShort'@'localhost';
create schema urlShort;
use urlShort;

create table shorts
(
	short text not null,
	`long` text not null
);

create unique index shorts_short_uindex
	on shorts (short);
