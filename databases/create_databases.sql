drop database if exists starter_test;
drop user starter;

create user starter with password 'starter';
create database starter_test;
grant all privileges on database starter_test to starter;
