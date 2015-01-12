require "sinatra"
require "sinatra/reloader" if development?
require "json"
#require "./database"

get "/" do
  send_file "views/home.html"
end

get "/permits" do
  content_type :json

  permit1 = {
     id: 1,
     created_at: DateTime.now,
     application_id: "123",
     file_date: DateTime.now,
     street_number: "123",
     street: "Duboce",
     street_suffix: "Av",
     city: "San Francisco",
  }

  permit2 = {
     id: 2,
     created_at: DateTime.now,
     application_id: "321",
     file_date: DateTime.now,
     street_number: "345",
     street: "Belcher",
     street_suffix: "St",
     city: "San Francisco",
  }

  [permit1, permit2].to_json
end
