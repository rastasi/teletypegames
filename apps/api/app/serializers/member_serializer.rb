class MemberSerializer
  def self.serialize_all(members)
    members.map { |m| serialize(m) }
  end

  def self.serialize(m)
    {
      nick:            m.nick,
      real_nick:       m.real_nick,
      motto:           m.motto,
      avatar_filename: m.avatar_filename,
      image_url:       m.image_id ? "/api/image/#{m.image_id}" : nil
    }
  end
end
