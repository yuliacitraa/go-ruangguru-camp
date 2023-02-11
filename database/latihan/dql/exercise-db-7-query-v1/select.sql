SELECT 
    id, 
    CONCAT(first_name, ' ', last_name) AS student_name, 
    student_class, 
    final_score, 
    absent 
    
    FROM reports 
    where final_score < 70 or absent > 5;