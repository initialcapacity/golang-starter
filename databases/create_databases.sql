drop database if exists starter_development;
drop database if exists starter_test;
drop user starter;

create database starter_development;
create database starter_test;
create user starter with password 'starter';
grant all privileges on database starter_development to starter;
grant all privileges on database starter_test to starter;
