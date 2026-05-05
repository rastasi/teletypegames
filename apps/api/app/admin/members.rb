ActiveAdmin.register Member do
  permit_params :nick, :real_nick, :motto, :avatar_filename

  menu priority: 4

  index do
    selectable_column
    id_column
    column :nick
    column :real_nick
    column :avatar_filename
    column :created_at
    actions
  end

  filter :nick
  filter :real_nick

  show do
    attributes_table do
      row :id
      row :nick
      row :real_nick
      row :motto
      row :avatar_filename
      row :created_at
      row :updated_at
    end
  end

  form do |f|
    f.inputs do
      f.input :nick
      f.input :real_nick
      f.input :motto
      f.input :avatar_filename
    end
    f.actions
  end
end
