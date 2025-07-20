class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line[/:\s+(.*?)\s*$/, 1]
  end

  def log_level
    @line[/\[(.*?)\]/, 1].downcase!
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
