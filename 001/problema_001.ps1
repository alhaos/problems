& {
  <#
    .SYNOPSIS
    Напишите программу, которая по введенному не более чем четырехзначному числу k будет выдавать сумму его цифр.
    
    .PARAMETER number
    На вход программе подается целое число k (0≤k≤9999).
 
    .OUTPUTS
    Выведите сумму его цифр.
 
#>

    param(
      [Parameter(Mandatory)]
      [ValidateRange(1, 9999)]
      [Int32]$Number
    )
    $result = [math]::Truncate($Number / 1000) % 10
    $result += [math]::Truncate($Number / 100) % 10
    $result += [math]::Truncate($Number / 10) % 10
    $result += [math]::Truncate($Number / 1) % 10
    return $result
  } 1234