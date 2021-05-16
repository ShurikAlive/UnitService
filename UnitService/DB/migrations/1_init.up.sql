CREATE TABLE `unit_db`.`units` (
    id varchar(50),
	Name varchar(150),
	ForceName varchar(50),
	Hp int,
	Initiative int,
	Bs int,
	Fs int,
	AdditionalRule varchar(500)
);

CREATE TABLE `unit_db`.`equipments` (
    id varchar(50),
	Name varchar(150),
	LimitOnUnit int,
	LimitOnTeam int,
	SoldarRole varchar(50),
	Rule varchar(500),
	Ammo int,
	Cost int
);