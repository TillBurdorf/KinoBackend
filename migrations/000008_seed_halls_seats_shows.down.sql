DELETE FROM shows WHERE hall_id = '6e429b44-1e9c-4dd5-bea9-b32baf770de9';

DELETE FROM seats WHERE hall_id IN ('6e429b44-1e9c-4dd5-bea9-b32baf770de9', '88e0431e-e124-45c8-8a13-a8d8a6b040da');

DELETE FROM halls WHERE id IN ('6e429b44-1e9c-4dd5-bea9-b32baf770de9', '88e0431e-e124-45c8-8a13-a8d8a6b040da')