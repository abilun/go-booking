CREATE KEYSPACE IF NOT EXISTS hotel WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TYPE IF NOT EXISTS hotel.address (
    country text,
    city text,
    street text,
    building smallint,
    entrance smallint,
    zip_code text
);

CREATE TABLE IF NOT EXISTS hotel.hotels (
    hotel_id uuid PRIMARY KEY,
    name text,
    address frozen<address>,
    description text,
    phone text,
    pois list<uuid>,
    rooms list<uuid>
);
