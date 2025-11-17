-- $1 points to label - code snippet part of golang database/sql package 
INSERT INTO TARGETGROUP (label) VALUES ($1) ON CONFLICT (label) DO NOTHING