ActiveAdmin.register_page "Dashboard" do
  menu priority: 1, label: proc { I18n.t("active_admin.dashboard") }

  content title: proc { I18n.t("active_admin.dashboard") } do
    columns do
      column do
        panel "Softwares" do
          ul do
            li "Total: #{Software.count}"
            li "Released: #{Software.where(status: 'released').count}"
            li "In development: #{Software.where(status: 'development').count}"
          end
        end
      end

      column do
        panel "Recent Releases" do
          table_for Release.order(created_at: :desc).limit(10) do
            column(:software) { |r| link_to r.software.title, admin_software_path(r.software) rescue r.software_id }
            column :version
            column :created_at
          end
        end
      end
    end
  end
end
