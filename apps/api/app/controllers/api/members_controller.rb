class Api::MembersController < ApiController
  def index
    members = Member.order(:id).map do |m|
      { nick: m.nick, real_nick: m.real_nick, motto: m.motto, avatar_filename: m.avatar_filename }
    end
    render json: members
  end
end
