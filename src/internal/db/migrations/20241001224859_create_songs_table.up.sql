CREATE TABLE IF NOT EXISTS songs (
                                     id SERIAL PRIMARY KEY,
                                     group_name TEXT NOT NULL,
                                     song TEXT NOT NULL,
                                     release_date DATE,
                                     text TEXT,
                                     link VARCHAR(255)
);

INSERT INTO songs (group_name, song, release_date, text, link) VALUES
                                                                        ('The Beatles', 'Hey Jude', '1968-08-26', 'Take a sad song and make it better.', 'http://example.com/hey_jude'),
                                                                        ('The Beatles', 'Let It Be', '1970-03-06', 'Whisper words of wisdom, let it be.', 'http://example.com/let_it_be'),
                                                                        ('Queen', 'Bohemian Rhapsody', '1975-10-31', 'Is this the real life? Is this just fantasy?', 'http://example.com/bohemian_rhapsody'),
                                                                        ('Queen', 'We Will Rock You', '1977-10-07', 'We will, we will rock you.', 'http://example.com/we_will_rock_you'),
                                                                        ('Nirvana', 'Smells Like Teen Spirit', '1991-09-10', 'With the lights out, it’s less dangerous.', 'http://example.com/smells_like_teen_spirit'),
                                                                        ('Nirvana', 'Come As You Are', '1992-03-02', 'Come as you are, as you were.', 'http://example.com/come_as_you_are'),
                                                                        ('Pink Floyd', 'Comfortably Numb', '1979-11-30', 'There is no pain, you are receding.', 'http://example.com/comfortably_numb'),
                                                                        ('Pink Floyd', 'Wish You Were Here', '1975-09-12', 'So, so you think you can tell.', 'http://example.com/wish_you_were_here'),
                                                                        ('Led Zeppelin', 'Stairway to Heaven', '1971-11-08', 'There’s a lady who’s sure all that glitters is gold\nAnd she is buying a stairway to heaven', 'http://example.com/stairway_to_heaven'),
                                                                        ('Led Zeppelin', 'Whole Lotta Love', '1969-08-07', 'You need coolin’, baby, I’m not foolin’.', 'http://example.com/whole_lotta_love');
