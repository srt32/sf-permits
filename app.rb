require "sinatra"
require "sinatra/reloader" if development?
require "json"
#require "./database"

get "/" do
  send_file "views/home.html"
end

get "/permits" do
  content_type :json

  permit = {
     id: 1,
     created_at: DateTime.now,
     application_id: "123",
     file_date: DateTime.now,
     street_number: "123",
     street: "Duboce",
     street_suffix: "Av",
     city: "San Francisco",
  }

  [permit].to_json
end
