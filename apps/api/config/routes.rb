Rails.application.routes.draw do
  devise_for :admin_users, ActiveAdmin::Devise.config
  ActiveAdmin.routes(self)

  namespace :api do
    get "software",             to: "software#index"
    get "software/highlighted", to: "software_highlighted#index"
    get "events",               to: "events#index"
    get "members",              to: "members#index"
    get "image/:id",            to: "images#show"
  end

  get "update",       to: "update#update"
  get "file(/*path)",   to: "files#show", format: false
end
