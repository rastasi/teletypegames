AdminUser.find_or_create_by!(email: ENV.fetch("ADMIN_EMAIL", "admin@teletype.hu")) do |user|
  user.password              = ENV.fetch("ADMIN_PASSWORD", "password123")
  user.password_confirmation = ENV.fetch("ADMIN_PASSWORD", "password123")
end

[
  { nick: "Mr. Zero",  real_nick: "Tasi",   motto: "His dream is to become an open-source knight.",                              avatar_filename: "mr.zero.png"  },
  { nick: "Mr. One",   real_nick: "Ballz",  motto: "The cheese is half-eaten. Something here went very wrong.",                  avatar_filename: "mr.one.png"   },
  { nick: "Mr. Two",   real_nick: "Z",      motto: "The egg was first or the chicken? It doesn't matter, we eat both.",          avatar_filename: "mr.two.png"   },
  { nick: "Mr. Three", real_nick: "gBird",  motto: "Life is beautiful because it contains the possibility of becoming human.",   avatar_filename: "mr.three.png" },
].each do |attrs|
  Member.find_or_create_by!(nick: attrs[:nick]) do |m|
    m.real_nick       = attrs[:real_nick]
    m.motto           = attrs[:motto]
    m.avatar_filename = attrs[:avatar_filename]
  end
end
