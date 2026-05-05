class Api::MembersController < ApiController
  def index
    members = Member.includes(:image).order(:id).map do |m|
      {
        nick:              m.nick,
        real_nick:         m.real_nick,
        motto:             m.motto,
        avatar_filename:   m.avatar_filename,
        image_url:         m.image_id ? "/api/image/#{m.image_id}" : nil
      }
    end
    render json: members
  end
end
