class Api::MembersController < ApiController
  def index
    render json: MemberSerializer.serialize_all(Member.includes(:image).order(:id))
  end
end
