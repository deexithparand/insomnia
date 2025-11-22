-- $1 points to label - code snippet part of golang database/sql package 
INSERT INTO TARGETGROUP (LABEL) 
VALUES 
    ($1)
ON CONFLICT (LABEL) DO NOTHING