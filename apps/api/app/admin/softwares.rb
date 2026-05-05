ActiveAdmin.register Software do
  permit_params :name, :title, :author, :desc, :story,
                :license, :platform, :status, :highlighted,
                external_links_attributes: [:id, :label, :url, :_destroy],
                releases_attributes: [:id, :_destroy]

  menu priority: 2

  index do
    selectable_column
    id_column
    column :name
    column :title
    column :platform
    column :status do |sw|
      status_tag sw.status
    end
    column :highlighted
    column :created_at
    actions
  end

  filter :name
  filter :title
  filter :author
  filter :platform, as: :select, collection: %w[tic80 ebitengine love]
  filter :status,   as: :select, collection: %w[development demo released archived]
  filter :highlighted

  show do
    attributes_table do
      row :id
      row :name
      row :title
      row :author
      row :platform
      row :status
      row :highlighted
      row :license
      row(:desc)  { |sw| simple_format sw.desc }
      row(:story) { |sw| simple_format sw.story }
      row :created_at
      row :updated_at
    end

    panel "External Links" do
      table_for resource.external_links do
        column :label
        column(:url) { |el| link_to el.url, el.url, target: "_blank" }
        column :created_at
        column do |el|
          link_to "Edit",   edit_admin_software_external_link_path(resource, el)
          concat " "
          link_to "Delete", admin_software_external_link_path(resource, el),
                  method: :delete, data: { confirm: "Are you sure?" }
        end
      end
      br
      link_to "Add External Link",
              new_admin_software_external_link_path(resource)
    end

    panel "Releases" do
      table_for resource.releases.order(created_at: :desc) do
        column :version
        column :html_folder_path
        column :cartridge_path
        column :source_path
        column :docs_folder_path
        column :created_at
      end
    end
  end

  form do |f|
    f.inputs "Details" do
      f.input :name
      f.input :title
      f.input :author
      f.input :platform, as: :select, collection: %w[tic80 ebitengine love]
      f.input :status,   as: :select, collection: %w[development demo released archived]
      f.input :highlighted
      f.input :license
      f.input :desc,  as: :text, input_html: { rows: 4 }
      f.input :story, as: :text, input_html: { rows: 8 }
    end

    f.inputs "External Links" do
      f.has_many :external_links, allow_destroy: true, new_record: true do |el|
        el.input :label
        el.input :url
      end
    end

    f.actions
  end
end
