select * from consultations order by patient_id;
select patient_id, date, diagnosis, content from consultations order by patient_id;

select * from patients;
select id as "patient_id",
       firstname,
       lastname,
       street,
       zip,
       city,
       birthdate,
       email,
       phonenumber1,
       phonenumber2,
       insurance,
       ssn
       from patients order by patient_id;

select patient_id, parent_id as "patient_id parent" from relationships order by patient_id;

select * from patients pat
    left join relationships on pat.id = relationships.patient_id
    left join patients as parents on parents.id = relationships.parent_id;


\copy (select * from patients order by id) to ../patients.csv with csv delimiter ',' header;