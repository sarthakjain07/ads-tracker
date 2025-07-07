-- Create ads table
CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    image_url TEXT NOT NULL,
    target_url TEXT NOT NULL
);

-- Create ad_clicks table
CREATE TABLE IF NOT EXISTS ad_clicks (
    id SERIAL PRIMARY KEY,
    ad_id INTEGER NOT NULL REFERENCES ads(id),
    timestamp TIMESTAMP NOT NULL,
    ip TEXT,
    playback_time INTEGER,
    UNIQUE (ad_id, timestamp, ip)
);

-- sample data
INSERT INTO ads (title, image_url, target_url) VALUES
('Ad 1', 'https://example.com/img1.jpg', 'https://target.com/1'),
('Ad 2', 'https://example.com/img2.jpg', 'https://target.com/2')
ON CONFLICT DO NOTHING;
