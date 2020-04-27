SELECT * FROM countries;

SELECT * FROM data
JOIN countries ON countries.id = data.country_id
WHERE countries.country = 'Australia' AND countries.state = 'New South Wales' AND date(data.date) = date('2020-04-25');