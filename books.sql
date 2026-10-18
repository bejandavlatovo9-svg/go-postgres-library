-- 1. Создание таблицы
CREATE TABLE booksss (
                         id SERIAL PRIMARY KEY,
                         title TEXT,
                         author TEXT,
                         price INT,
                         page INT,
                         created_at TIMESTAMP DEFAULT NOW()
);

-- 2. Добавление 10 книг
INSERT INTO booksss (title, author, price, page)
VALUES
    ('C++ Basics', 'Ali Rahimov', 200, 250),
    ('Python Beginner', 'Bobur Saidov', 300, 320),
    ('JavaScript Guide', 'Yusuf Karimov', 400, 400),
    ('Go Programming', 'Farid Sharipov', 450, 350),
    ('SQL for Beginners', 'Aziz Davlatov', 250, 280),
    ('HTML and CSS', 'Malika Sobirova', 180, 200),
    ('Java Fundamentals', 'Said Naimov', 500, 500),
    ('Learn PostgreSQL', 'Rustam Aliyev', 350, 310),
    ('Algorithms', 'Kamol Nazarov', 600, 450),
    ('Computer Science', 'Zarina Olimova', 220, 300);

-- 3. Вывести все книги
SELECT * FROM booksss;

-- 4. Вывести название и автора
SELECT title, author FROM booksss;

-- 5. Найти книгу с id = 5
SELECT * FROM booksss
WHERE id = 5;

-- 6. Изменить цену книги с id = 3
UPDATE booksss
SET price = 350
WHERE id = 3;

-- 7. Увеличить цену всех книг на 100
UPDATE booksss
SET price = price + 100;

-- 8. Изменить страницы книги с id = 6
UPDATE booksss
SET page = 500
WHERE id = 6;

-- 9. Удалить книгу с id = 8
DELETE FROM booksss
WHERE id = 8;

-- 10. Удалить книги дешевле 250
DELETE FROM booksss
WHERE price < 250;

-- 11. Вывести книги дороже 400
SELECT * FROM booksss
WHERE price > 400;

-- 12. Вывести книги, у которых страниц больше 300
SELECT * FROM booksss
WHERE page > 300;

-- 13. Сортировка цены от большей к меньшей
SELECT * FROM booksss
ORDER BY price DESC;

-- 14. Вывести первые 5 книг
SELECT * FROM booksss
ORDER BY id
LIMIT 5;
SELECT current_database(), current_schema();
