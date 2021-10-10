SELECT u.ID, u.UserName, p.UserName as ParentUserName 
  FROM USER as u
  LEFT JOIN USER as p on u.Parent = p.ID;