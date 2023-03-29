task :default => :git

task :git do
  puts "请输入git内容 m"
    @msg = STDIN.gets.chomp

    @slug = "#{@msg}"
    @slug = @slug.downcase.strip.gsub(' ', '-')
    @date = Time.now.strftime("%F %T")
    @finalmsg = "feat: #{@slug} #{@date}"
    system "git add ."
    system "git commit -m \"#{@finalmsg}\""
    system "git status"
    system "git push"
    # system "git push github"
    # system "git push gitee"
    puts "git push \"#{@finalmsg}\""
end