STUDENT TABLE:
ID - primary key
name - varchar
email - varchar
standard - varchar
motherName - varchar
fatherName - varchar
address - varchar
mobile - varchar
dateofBirth - timestamp
dateOfJoining - timestamp
createdAt - timestamp
modifiedAt - timestamp


AUTH TABLE:

ID - primary key
studentID - varchar (from STUDENT table)
password - varchar
createdAt - timestamp
modifiedAt - timestamp


ATTENDANCE_ENTRY TABLE:

ID - primary key
studentID - varchar (from STUDENT table) (It may duplicate because of multiple entry)
date - timestamp
checkIn - timestamp
checkout - timestamp

WORKDONE_ENTRY TABLE:

ID - primary key
studentID - varchar (from STUDENT table) (It may duplicate because of multiple entry)
date - timestamp
workDone - varchar


LEAVE_ENTRY TABLE:

ID - primary key
studentID - varchar (from STUDENT table) (It may duplicate because of multiple entry)
date - timestamp
reason - varchar
