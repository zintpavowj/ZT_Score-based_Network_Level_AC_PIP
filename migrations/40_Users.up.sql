CREATE TABLE IF NOT EXISTS users (
	id                        SERIAL PRIMARY KEY,
	name                      VARCHAR(100) UNIQUE NOT NULL,
    email                     VARCHAR(100) UNIQUE NOT NULL,
    last_access_time          TIMESTAMPTZ NOT NULL,
    expected                  FLOAT NOT NULL,
    access_time_min           TIME NOT NULL,
    access_time_max           TIME NOT NULL,
    database_update_time      TIMESTAMPTZ NOT NULL,
    password_failed_attempts  INT NOT NULL
);

INSERT INTO 
    users (name, email, last_access_time, expected, access_time_min, access_time_max, database_update_time, password_failed_attempts)
VALUES
    ('Clare Fuller', 'clarefuller@hotmail.couk', '2023-03-06 06:47:28 +1:00', '0.911086408913606', '07:14', '07:56', '2023-03-06 06:47:28 +1:00', 3),
    ('Cassidy Mckee', 'cassidymckee@hotmail.net', '2023-03-06 05:47:19 +1:00', '0.691370132474801', '04:26', '11:09', '2023-03-06 05:47:19 +1:00', 3),
    ('Cathleen Parsons', 'cathleenparsons@outlook.net', '2023-03-05 19:28:01 +1:00', '0.644571503565047', '03:11', '05:59', '2023-03-05 19:28:01 +1:00', 0),
    ('Dylan Bender', 'dylanbender1666@yahoo.org', '2023-03-05 14:07:46 +1:00', '0.265379698950691', '04:02', '08:11', '2023-03-05 14:07:46 +1:00', 0),
    ('Gisela Ellis', 'giselaellis3687@aol.com', '2023-03-06 03:44:38 +1:00', '0.480675778710478', '06:12', '01:58', '2023-03-06 03:44:38 +1:00', 1),
    ('Pearl Ortiz', 'pearlortiz4929@outlook.couk', '2023-03-06 02:20:39 +1:00', '0.193931733946456', '04:25', '05:32', '2023-03-06 02:20:39 +1:00', 1),
    ('Nathan Daniel', 'nathandaniel8312@yahoo.couk', '2023-03-05 12:48:10 +1:00', '0.0075812274671887', '05:53', '07:23', '2023-03-05 12:48:10 +1:00', 4),
    ('Alfreda Blevins', 'alfredablevins4430@outlook.edu', '2023-03-05 23:40:34 +1:00', '0.221253517936135', '11:38', '03:07', '2023-03-05 23:40:34 +1:00', 4),
    ('Zahir Hicks', 'zahirhicks4238@icloud.ca', '2023-03-06 03:09:42 +1:00', '0.236651743437306', '10:47', '04:17', '2023-03-06 03:09:42 +1:00', 0),
    ('Unity Booth', 'unitybooth42@outlook.couk', '2023-03-06 01:18:11 +1:00', '0.931146366009627', '06:32', '05:18', '2023-03-06 01:18:11 +1:00', 0);
