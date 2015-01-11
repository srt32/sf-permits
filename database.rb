require "sinatra/sequel"

set :database, (ENV["DATABASE_URL"] || "postgres://localhost/permits_development")

migration "create permits table" do
  database.create_table :permits do
    primary_key :id
    timestamp :created_at, null: false

    integer :application_id
    timestamp :file_date
    text :street_number
    text :street
    text :street_suffix
    text :city
  end
end
