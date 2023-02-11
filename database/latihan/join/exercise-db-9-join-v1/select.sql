select
    reports.id as id,
    students.fullname as fullname,
    students.class as class,
    students.status as status,
    reports.study as study,
    reports.score as score

    from reports
    join students on reports.student_id = students.id
    where students.status='active' and reports.score<70 order by reports.score asc;