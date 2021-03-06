import * as React from "react";
import styled from "styled-components";

import { PostInterface } from "../../types";
import Button from "../../components/Button";
import FavouriteButton from "../../components/FavouriteButton";
import {
  ContentContainer,
  Avatar,
  AvatarImage,
  Info,
  InfoHeader,
  Name,
  Headline,
  Actions,
  Title,
  Description,
} from "./PostComponents";
import { Menu, StyledMenuList, StyledMenuItem } from "../../components/Menu";

const Container = styled.div`
  background-color: white;
  border-radius: 0.875rem;
  padding-left: 0.75rem;
`;

const ButtonContainer = styled.div`
  padding: 1rem 1.5rem 1.75rem 0.75rem;
`;

type PostDetailProps = {
  data: PostInterface;
};

const EditIcon = () => (
  <svg width="23" height="23" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
      d="M10 4H3a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7"
      stroke="#23374D"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
    <path
      d="M17.5 2.5a2.121 2.121 0 113 3L11 15l-4 1 1-4 9.5-9.5z"
      stroke="#23374D"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    />
  </svg>
);

const DeleteIcon = () => (
  <svg width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path
      d="M8.438 4.313H8.25a.188.188 0 00.188-.188v.188h7.124v-.188c0 .103.085.188.188.188h-.188V6h1.688V4.125c0-.827-.673-1.5-1.5-1.5h-7.5c-.827 0-1.5.673-1.5 1.5V6h1.688V4.312zM20.25 6H3.75a.75.75 0 00-.75.75v.75c0 .103.084.188.188.188h1.415l.579 12.257c.038.8.698 1.43 1.498 1.43h10.64c.802 0 1.46-.628 1.498-1.43l.579-12.258h1.416A.188.188 0 0021 7.5v-.75a.75.75 0 00-.75-.75zm-3.11 13.688H6.86l-.567-12h11.414l-.567 12z"
      fill="#FE4A4A"
    />
  </svg>
);

const DangerText = styled.span`
  color: var(--color-red);
`;

const PostDetail: React.FC<PostDetailProps> = function (props) {
  const { data } = props;
  const { created_by: user } = data;
  return (
    <Container>
      <ContentContainer>
        <Avatar>
          <AvatarImage src={user.avatar_url} />
        </Avatar>
        <Info>
          <InfoHeader>
            <div>
              <Name>
                {user.first_name} {user.last_name}
              </Name>
              <Headline>
                {user.job_title} at {user.company}
              </Headline>
            </div>
            <Actions>
              <FavouriteButton active={data.favourite} />
              <Menu>
                <StyledMenuList>
                  <StyledMenuItem>
                    <EditIcon />
                    Edit
                  </StyledMenuItem>
                  <StyledMenuItem>
                    <DeleteIcon />
                    <DangerText>Delete Post</DangerText>
                  </StyledMenuItem>
                </StyledMenuList>
              </Menu>
            </Actions>
          </InfoHeader>
          <Title>{data.title}</Title>
          <Description>{data.description}</Description>
        </Info>
      </ContentContainer>
      <ButtonContainer>
        <Button fullWidth primary>
          Connect with {user.first_name}
        </Button>
      </ButtonContainer>
    </Container>
  );
};

export default PostDetail;
