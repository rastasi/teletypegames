ActiveAdmin.register Image do
  permit_params :file_upload

  menu priority: 5

  index do
    selectable_column
    id_column
    column :original_filename
    column :content_type
    column(:preview) do |img|
      if File.exist?(img.file_path)
        image_tag("/api/image/#{img.id}", style: "max-height:60px;max-width:120px;object-fit:contain;")
      end
    end
    column :created_at
    actions
  end

  filter :original_filename
  filter :content_type

  show do
    attributes_table do
      row :id
      row :original_filename
      row :filename
      row :content_type
      row(:preview) do |img|
        if File.exist?(img.file_path)
          image_tag("/api/image/#{img.id}", style: "max-height:300px;max-width:100%;object-fit:contain;")
        else
          "File not found on disk"
        end
      end
      row :created_at
      row :updated_at
    end
  end

  form html: { multipart: true } do |f|
    f.inputs do
      f.input :file_upload, as: :file, label: "Image File"
    end
    f.actions
  end
end
