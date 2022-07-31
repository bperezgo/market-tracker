CREATE TABLE dummy_asset (
    id UUID UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    price REAL NOT NULL
);