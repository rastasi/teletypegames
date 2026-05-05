ActiveAdmin.register Event do
  permit_params :name, :date

  menu priority: 3

  index do
    selectable_column
    id_column
    column :name
    column :date
    column :created_at
    actions
  end

  filter :name
  filter :date

  show do
    attributes_table do
      row :id
      row :name
      row :date
      row :created_at
      row :updated_at
    end
  end

  form do |f|
    f.inputs do
      f.input :name
      f.input :date, as: :date_time_picker
    end
    f.actions
  end
end
