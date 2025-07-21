module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city[...4].upcase.to_sym
  end

  TERMINAL_A = :A
  TERMINAL_B = :B
  
  def self.get_terminal(ship_identifier)
    case ship_identifier.to_s[...3].upcase
      in "OIL" | "GAS"
        TERMINAL_A
      else
        TERMINAL_B
    end
  end
end
