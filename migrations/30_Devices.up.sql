CREATE TABLE IF NOT EXISTS devices (
	id               SERIAL PRIMARY KEY,
	name             VARCHAR(100) NOT NULL,
    cert_cn          VARCHAR(255) UNIQUE NOT NULL,
    last_access_time TIMESTAMPTZ NOT NULL,
    expected         FLOAT NOT NULL
);

INSERT INTO 
    devices (name, cert_cn, last_access_time, expected)
VALUES
    ('Notebook of Clare Fuller', 'Notebook_ClaFul', '2023-03-07 09:12:19  +1:00', 0.869355948839507),
    ('Notebook of Cassidy Mckee', 'Notebook_CasMck', '2023-03-06 20:58:51  +1:00', 0.781210265691283),
    ('Notebook of Cathleen Parsons', 'Notebook_CatPar', '2023-03-07 05:50:39  +1:00', 0.513965860731749),
    ('Notebook of Dylan Bender', 'Notebook_DylBen', '2023-03-07 06:50:44  +1:00', 0.904638875681582),
    ('Notebook of Gisela Ellis', 'Notebook_GisEll', '2023-03-07 12:59:34  +1:00', 0.40683793684127),
    ('Notebook of Pearl Ortiz', 'Notebook_PeaOrt', '2023-03-07 12:49:18  +1:00', 0.828342746858796),
    ('Notebook of Nathan Daniel', 'Notebook_NatDan', '2023-03-06 16:14:39  +1:00', 0.43850375946819),
    ('Notebook of Alfreda Blevins', 'Notebook_AlfBle', '2023-03-07 01:16:36  +1:00', 0.135498456313309),
    ('Notebook of Zahir Hicks', 'Notebook_ZahHic', '2023-03-06 21:25:32  +1:00', 0.0460150957788441),
    ('Notebook of Unity Booth', 'Notebook_UniBoo', '2023-03-06 22:50:18  +1:00', 0.51228326469137),
    ('iPhone of Clare Fuller', 'iPhone_ClaFul', '2023-03-06 19:28:39  +1:00', 0.00574209525212799),
    ('iPhone of Cassidy Mckee', 'iPhone_CasMck', '2023-03-06 21:05:49  +1:00', 0.633742414521277),
    ('iPhone of Cathleen Parsons', 'iPhone_CatPar', '2023-03-06 16:23:09  +1:00', 0.889514974308707),
    ('iPhone of Dylan Bender', 'iPhone_DylBen', '2023-03-07 05:32:36  +1:00', 0.0761677081160872),
    ('iPhone of Gisela Ellis', 'iPhone_GisEll', '2023-03-06 17:15:49  +1:00', 0.709290991221581),
    ('iPhone of Pearl Ortiz', 'iPhone_PeaOrt', '2023-03-07 06:57:10  +1:00', 0.730371029496333),
    ('Phone of Dylan Bender', 'Phone_DylBen', '2023-03-06 15:36:05  +1:00', 0.453265596227542),
    ('Phone of Gisela Ellis', 'Phone_GisEll', '2023-03-07 03:39:58  +1:00', 0.997007658409745),
    ('Phone of Pearl Ortiz', 'Phone_PeaOrt', '2023-03-07 08:47:27  +1:00', 0.199259321735573),
    ('Phone of Nathan Daniel', 'Phone_NatDan', '2023-03-07 11:20:06  +1:00', 0.272394398180995),
    ('Phone of Alfreda Blevins', 'Phone_AlfBle', '2023-03-07 11:34:34  +1:00', 0.184949898631211),
    ('Phone of Zahir Hicks', 'Phone_ZahHic', '2023-03-07 11:59:28  +1:00', 0.161805532797823),
    ('Phone of Unity Booth', 'Phone_UniBoo', '2023-03-07 01:20:59  +1:00', 0.331808504521491),
    ('Tablet of Clare Fuller', 'Tablet_ClaFul', '2023-03-06 20:29:39  +1:00', 0.634367745433711),
    ('Tablet of Cassidy Mckee', 'Tablet_CasMck', '2023-03-06 19:24:00  +1:00', 0.370517981866762),
    ('Tablet of Dylan Bender', 'Tablet_DylBen', '2023-03-07 03:54:31  +1:00', 0.336440373684416),
    ('Tablet of Gisela Ellis', 'Tablet_GisEll', '2023-03-07 09:37:12  +1:00', 0.426078277055711),
    ('Tablet of Alfreda Blevins', 'Tablet_AlfBle', '2023-03-07 10:38:04  +1:00', 0.566373843819547);
