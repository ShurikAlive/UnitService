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

INSERT INTO `unit_db`.`equipments`
(`id`,`Name`,`LimitOnUnit`,`LimitOnTeam`,`SoldarRole`,`Rule`,`Ammo`,`Cost`)
VALUES
('780fa83b-dacf-4d84-b0a8-13620ed57b35','M230',1,1,'Grenader','2 orders: The under barrel grenade launcher deals 1 wound to all models in its AoE. The target must be in LoS and does not have to be a model. No max range to this tossed weapon.',2,3)
,('0abf8d9c-8c8c-4706-9c97-c85121e3545f','Medic',1,1,'Medic','This soldier is the fireteam is medic. He does not roll to perform medical aid on other soldiers (he must still use 1 order). When this soldier performs medical aid to another soldier that has suffered 3 wounds, the -1 modafire for receiving 3 wounds is removed. This soldier cannot benefit from Medic on himself.',-1,3)
,('5131d405-88fc-4ee8-aa1b-f92da9dbe883','Machine gunner M249 LMG',1,1,'Machine gunner','This soldier is the fireteam is gunner. The M249 replaces the M4 carbine as his primary weapon. This soldier can perform Suppression Fire for 1 order. This weapon deals 2 wounds. This soldier receives a -2 to his accuracy.',-1,3)
,('b9055034-e631-4e65-b2e9-40ecaf81115b','Team Leader',1,1,'Team Leader','This soldier is the fireteam is team leader. He receives +1 to his initiative roll.',-1,3)
,('b1256ae9-b91e-4ef0-8283-87c1687a6c4b','M67 Grenade',-1,-1,'','1 order: This tossed weapon deals 1 wounds to all models in its AoE.',1,2)
,('0eddeb73-801d-4e58-b795-a1334d5d2879','M84 Stun Grenade',-1,-1,'','1 orders: This tossed weapon renders all models in its AoE inactive',2,2);


INSERT INTO `unit_db`.`units`
(`id`,`Name`,`ForceName`,`Hp`,`Initiative`,`Bs`,`Fs`,`AdditionalRule`)
VALUES
('4c2732e8-522c-4cda-80b2-2a7762eb4ee2','Robert A. Pepper','Special Forces Airborne',4,8,1,4,'')
,('83b1a1e7-0927-459c-905f-57789f9cf2ec','Robert R. River','Special Forces Airborne',4,8,3,2,'')
,('780b1856-951b-4ed6-a93d-4a3cd16a93b9','Patrick Singletary','Special Forces Airborne',4,6,2,2,'')
,('aa6f736f-4dfd-4e06-b16c-3bccd2d99fd9','Gregiory W. Morris','Special Forces Airborne',4,8,1,1,'When Morris moves into hand to hand combat his roll is at +2')
,('edfab252-77c0-4001-8843-975129a67e33','John P. Crowder','Special Forces Airborne',4,6,1,2,'')
,('1b0c3709-fd65-459a-8d8b-138c792795e2','Stevphen C. Baybrook','Special Forces Airborne',4,5,1,1,'');






















