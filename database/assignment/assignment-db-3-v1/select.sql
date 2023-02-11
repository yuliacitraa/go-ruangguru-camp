

SELECT 
    id, 
    CONCAT(first_name, ' ', last_name) AS fullname, 
    substring(exam_id, 1, 2) AS class, 
    AVG( bahasa_indonesia, bahasa_inggris, matematika, ipa) AS average_score  
    
    FROM final_scores
    WHERE exam_status = 'pass' AND fee_status in('full', 'installment')
    ORDER BY average_score DESC LIMIT 5;