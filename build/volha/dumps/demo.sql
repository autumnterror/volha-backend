-- demo_seed.sql (xid ids)
BEGIN;

-- ========== brands ==========
INSERT INTO brands (id, title) VALUES
('d506a6nbne320j90g2kg', 'Nordic Atelier'),
('d506a6nbne320j90g2l0', 'Baltic Form'),
('d506a6nbne320j90g2lg', 'Arctic Line');

-- ========== categories ==========
INSERT INTO categories (id, title, uri, img) VALUES
('d506a6nbne320j90g2m0', 'Кресла', 'chairs', '1.jpg'),
('d506a6nbne320j90g2mg', 'Столы',  'tables', '2.jpg'),
('d506a6nbne320j90g2n0', 'Свет',   'lights', '3.jpg');

-- ========== countries ==========
INSERT INTO countries (id, title, friendly) VALUES
('d506a6nbne320j90g2ng', 'Finland', 'Финляндия'),
('d506a6nbne320j90g2o0', 'Estonia', 'Эстония'),
('d506a6nbne320j90g2og', 'Sweden',  'Швеция');

-- ========== materials ==========
INSERT INTO materials (id, title) VALUES
('d506a6nbne320j90g2p0', 'Массив дуба'),
('d506a6nbne320j90g2pg', 'Сталь порошковая'),
('d506a6nbne320j90g2q0', 'Лён премиум');

-- ========== colors ==========
INSERT INTO colors (id, title, hex) VALUES
('d506a6nbne320j90g2qg', 'Снег',   '#F7F7F7'),
('d506a6nbne320j90g2r0', 'Графит', '#2B2B2B'),
('d506a6nbne320j90g2rg', 'Мох',    '#3A5F3A');

-- ========== products ==========
INSERT INTO products (
  id, title, article, brand_id, category_id, country_id,
  width, height, depth, photos, price, description, views, is_favorite
) VALUES
(
  'd506a6nbne320j90g2s0',
  'Кресло Fjord Soft',
  'ART-FJORD-001',
  'd506a6nbne320j90g2kg',
  'd506a6nbne320j90g2m0',
  'd506a6nbne320j90g2ng',
  780, 920, 820,
  ARRAY['1.jpg','2.jpg','3.jpg'],
  25900,
  'Мягкое кресло с выраженной поддержкой спины. Подходит для гостиной и кабинета.',
  1000,
  false
),
(
  'd506a6nbne320j90g2sg',
  'Стол Baltic Oak 140',
  'ART-BALTIC-140',
  'd506a6nbne320j90g2l0',
  'd506a6nbne320j90g2mg',
  'd506a6nbne320j90g2o0',
  1400, 760, 800,
  ARRAY['2.jpg','3.jpg','1.jpg'],
  39900,
  'Обеденный стол на металлическом основании. Столешница из дуба, устойчивое покрытие.',
  500,
  false
),
(
  'd506a6nbne320j90g2t0',
  'Светильник Aurora Cone',
  'ART-AURORA-010',
  'd506a6nbne320j90g2lg',
  'd506a6nbne320j90g2n0',
  'd506a6nbne320j90g2og',
  260, 380, 260,
  ARRAY['3.jpg','1.jpg','2.jpg'],
  12900,
  'Подвесной светильник с мягким рассеиванием. Хорошо смотрится над столом или барной стойкой.',
  5000,
  true
);

-- ========== product_materials ==========
INSERT INTO product_materials (product_id, material_id) VALUES
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2q0'),
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2pg'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2p0'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2pg'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2pg'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2p0');

-- ========== product_colors ==========
INSERT INTO product_colors (product_id, color_id) VALUES
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2qg'),
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2r0'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2r0'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2rg'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2qg'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2rg');

-- ========== product_seems ==========
INSERT INTO product_seems (product_id, similar_product_id) VALUES
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2sg'),
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2t0'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2s0'),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2t0'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2s0'),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2sg');

-- ========== product_color_photos ==========
INSERT INTO product_color_photos (product_id, color_id, photos) VALUES
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2qg', ARRAY['1.jpg','2.jpg']),
('d506a6nbne320j90g2s0', 'd506a6nbne320j90g2r0', ARRAY['2.jpg','3.jpg']),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2r0', ARRAY['2.jpg','1.jpg']),
('d506a6nbne320j90g2sg', 'd506a6nbne320j90g2rg', ARRAY['3.jpg','2.jpg']),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2qg', ARRAY['1.jpg','3.jpg']),
('d506a6nbne320j90g2t0', 'd506a6nbne320j90g2rg', ARRAY['3.jpg','1.jpg']);

-- ========== slides ==========
INSERT INTO slides (id, link, img, img762) VALUES
('d506a6nbne320j90g2tg', '/catalog/chairs', '1.jpg', '1.jpg'),
('d506a6nbne320j90g2u0', '/catalog/tables', '2.jpg', '2.jpg'),
('d506a6nbne320j90g2ug', '/catalog/lights', '3.jpg', '3.jpg');

COMMIT;
