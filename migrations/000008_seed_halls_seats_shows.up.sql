INSERT INTO halls (id, name, capacity) VALUES
('6e429b44-1e9c-4dd5-bea9-b32baf770de9', 'Kinsoaal 1', 7),
('88e0431e-e124-45c8-8a13-a8d8a6b040da', 'Kinosaal 2', 3);

INSERT INTO seats (id, hall_id, seat_row, seat_number) VALUES
('2b059ad6-b929-4579-98a5-c83fd40d2889', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 1, 1),
('5065903c-6587-414b-88ab-1cda0b77d379', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 1, 2),
('7a7d5bbe-299f-4408-9d36-0e4229ca69a5', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 1, 3),
('15e3868a-d421-4edb-b73a-95426aa08bd6', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 2, 1),
('2cd2fe26-2fda-4598-b0be-69790abec4d1', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 2, 2),
('1ed23f1c-f05f-407f-bed5-ca398fcef59e', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 2, 3),
('0a2d27ef-15b5-47c5-a9cf-8f23a4a9056d', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', 2, 4),

('ce6d87fe-274c-478a-9877-cbfa97b4119a', '88e0431e-e124-45c8-8a13-a8d8a6b040da', 1, 1),
('f007ae51-31e4-49f7-b717-0b6f0ab6bf07', '88e0431e-e124-45c8-8a13-a8d8a6b040da', 1, 2),
('3209ba9b-13b8-4a45-a6bc-f76791899bce', '88e0431e-e124-45c8-8a13-a8d8a6b040da', 2, 1);

INSERT INTO shows (id, movie_id, hall_id, start_time, buffer_minutes) VALUES
('9df90ba6-5ada-41c0-9e63-140f5b0ac97e', 'f053a164-c540-48cb-9641-7cd8a050307d', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', '2025-09-20 18:00:00', 15),
('51a7c573-1465-44b2-9a32-c57a8a7d816e', 'f053a164-c540-48cb-9641-7cd8a050307d', '6e429b44-1e9c-4dd5-bea9-b32baf770de9', '2025-09-20 20:15:00', 15);