ActiveAdmin.register Release do
  actions :index, :show

  menu false

  belongs_to :software, optional: true

  index do
    selectable_column
    id_column
    column(:software) { |r| link_to r.software.title, admin_software_path(r.software) }
    column :version
    column :html_folder_path
    column :created_at
    actions only: [:show]
  end

  filter :version

  show do
    attributes_table do
      row :id
      row(:software) { |r| link_to r.software.title, admin_software_path(r.software) }
      row :version
      row :cartridge_path
      row :source_path
      row :html_folder_path
      row :docs_folder_path
      row :created_at
      row :updated_at
    end
  end
end
