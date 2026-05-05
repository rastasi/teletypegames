ActiveAdmin.register Software do
  permit_params :name, :title, :author, :desc, :story,
                :license, :platform, :status, :highlighted, :image_id,
                external_links_attributes: [:id, :label, :url, :_destroy],
                releases_attributes: [:id, :_destroy]

  menu priority: 2

  index do
    selectable_column
    id_column
    column :name do |sw|
      link_to sw.name, admin_software_path(sw)
    end
    column :title
    column :platform
    column :status do |sw|
      status_tag sw.status
    end
    column :highlighted
    column(:image) do |sw|
      if sw.image && File.exist?(sw.image.file_path)
        image_tag "/api/image/#{sw.image.id}", style: "max-height:40px;max-width:80px;object-fit:contain;"
      end
    end
    column :created_at
    actions
  end

  filter :name
  filter :title
  filter :author
  filter :platform, as: :select, collection: %w[tic80 ebitengine love]
  filter :status,   as: :select, collection: %w[development demo released archived]
  filter :highlighted

  show title: proc { |sw| sw.title } do
    active_admin_form_for [:admin, resource], url: admin_software_path(resource), html: { method: :put } do |f|
      f.inputs "Details" do
        f.input :name
        f.input :title
        f.input :author
        f.input :platform, as: :select, collection: %w[tic80 ebitengine love]
        f.input :status,   as: :select, collection: %w[development demo released archived]
        f.input :highlighted
        f.input :license
        f.input :image_id, as: :select, label: "Image",
                           collection: Image.order(:original_filename).map { |img| [img.original_filename, img.id] },
                           include_blank: "— no image —"
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
      f.input :image_id, as: :select, label: "Image",
                         collection: Image.order(:original_filename).map { |img| [img.original_filename, img.id] },
                         include_blank: "— no image —"
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
