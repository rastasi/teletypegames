ActiveAdmin.register Member do
  permit_params :nick, :real_nick, :motto, :avatar_filename, :image_id

  menu priority: 4

  index do
    selectable_column
    id_column
    column :nick do |m|
      link_to m.nick, edit_admin_member_path(m)
    end
    column :real_nick
    column :motto do |m|
      truncate m.motto, length: 60
    end
    column :avatar_filename
    column(:image) do |m|
      if m.image && File.exist?(m.image.file_path)
        image_tag "/api/image/#{m.image.id}", style: "max-height:40px;max-width:80px;object-fit:contain;"
      end
    end
    column :created_at
    actions
  end

  filter :nick
  filter :real_nick

  show title: proc { |m| m.nick } do
    active_admin_form_for [:admin, resource], url: admin_member_path(resource), html: { method: :put } do |f|
      f.inputs "Edit Member" do
        f.input :nick
        f.input :real_nick
        f.input :motto
        f.input :avatar_filename,
                hint: "Fallback image filename (e.g. mr.zero.png) — only used when no image is selected above."
        f.input :image_id, as: :select, label: "Image",
                        collection: Image.order(:original_filename).map { |img| [img.original_filename, img.id] },
                        include_blank: "— no image (use avatar_filename) —"
      end
      f.actions
    end
  end

  form do |f|
    f.inputs do
      f.input :nick
      f.input :real_nick
      f.input :motto
      f.input :avatar_filename,
              hint: "Fallback image filename (e.g. mr.zero.png) — only used when no image is selected above."
      f.input :image_id, as: :select, label: "Image",
                      collection: Image.order(:original_filename).map { |img| [img.original_filename, img.id] },
                      include_blank: "— no image (use avatar_filename) —"
    end
    f.actions
  end
end
